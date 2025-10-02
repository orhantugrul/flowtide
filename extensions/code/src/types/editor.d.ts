export interface Editor {
  id?: number;
  name: string;
  version: string;
  createdAt?: Date;
  updatedAt?: Date;
  deletedAt?: Date;
}

export interface EditorQueryInput {
  name: string;
  version: string;
}

export interface EditorCreateInput {
  name: string;
  version: string;
}
