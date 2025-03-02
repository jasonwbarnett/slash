import { create } from "zustand";
import { collectionServiceClient } from "@/grpcweb";
import { Collection } from "@/types/proto/api/v2/collection_service";

interface CollectionState {
  collectionMapById: Record<number, Collection>;
  fetchCollectionList: () => Promise<Collection[]>;
  getOrFetchCollectionById: (id: number) => Promise<Collection>;
  getCollectionById: (id: number) => Collection;
  getCollectionList: () => Collection[];
  fetchCollectionByName: (collectionName: string) => Promise<Collection>;
  createCollection: (collection: Collection) => Promise<Collection>;
  updateCollection: (collection: Partial<Collection>, updateMask: string[]) => Promise<Collection>;
  deleteCollection: (id: number) => Promise<void>;
}

const useCollectionStore = create<CollectionState>()((set, get) => ({
  collectionMapById: {},
  fetchCollectionList: async () => {
    const { collections } = await collectionServiceClient.listCollections({});
    const collectionMap = get().collectionMapById;
    collections.forEach((collection) => {
      collectionMap[collection.id] = collection;
    });
    set(collectionMap);
    return collections;
  },
  getOrFetchCollectionById: async (id: number) => {
    const collectionMap = get().collectionMapById;
    if (collectionMap[id]) {
      return collectionMap[id] as Collection;
    }

    const { collection } = await collectionServiceClient.getCollection({
      id: id,
    });
    if (!collection) {
      throw new Error(`Collection with id ${id} not found`);
    }

    collectionMap[id] = collection;
    set(collectionMap);
    return collection;
  },
  getCollectionById: (id: number) => {
    const collectionMap = get().collectionMapById;
    return collectionMap[id] as Collection;
  },
  getCollectionList: () => {
    return Object.values(get().collectionMapById);
  },
  fetchCollectionByName: async (collectionName: string) => {
    const { collection } = await collectionServiceClient.getCollectionByName({
      name: collectionName,
    });
    if (!collection) {
      throw new Error(`Collection with name ${collectionName} not found`);
    }

    const collectionMap = get().collectionMapById;
    collectionMap[collection.id] = collection;
    set(collectionMap);
    return collection;
  },
  createCollection: async (collection: Collection) => {
    const { collection: createdCollection } = await collectionServiceClient.createCollection({
      collection: collection,
    });
    if (!createdCollection) {
      throw new Error(`Failed to create collection`);
    }

    const collectionMap = get().collectionMapById;
    collectionMap[createdCollection.id] = createdCollection;
    set(collectionMap);
    return createdCollection;
  },
  updateCollection: async (collection: Partial<Collection>, updateMask: string[]) => {
    const { collection: updatedCollection } = await collectionServiceClient.updateCollection({
      collection: collection,
      updateMask: updateMask,
    });
    if (!updatedCollection) {
      throw new Error("Collection not found");
    }

    const collectionMap = get().collectionMapById;
    collectionMap[updatedCollection.id] = updatedCollection;
    set(collectionMap);
    return updatedCollection;
  },
  deleteCollection: async (id: number) => {
    await collectionServiceClient.deleteCollection({
      id,
    });
    const collectionMap = get().collectionMapById;
    delete collectionMap[id];
    set(collectionMap);
  },
}));

export default useCollectionStore;
