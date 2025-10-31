export interface Editor {
  id: number;
  name: string;
  version: string;
  createdAt?: string;
  updatedAt?: string;
  deletedAt?: string;
}

export interface EditorQueryInput {
  name: string;
  version: string;
}

export interface EditorCreateInput {
  name: string;
  version: string;
}
