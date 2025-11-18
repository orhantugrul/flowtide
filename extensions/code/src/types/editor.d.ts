export type Editor = {
  id: number;
  name: string;
  version: string;
  createdAt?: string;
  updatedAt?: string;
  deletedAt?: string;
};

export type EditorQueryInput = {
  name?: string;
  version?: string;
};

export type EditorCreateInput = {
  name: string;
  version: string;
};
