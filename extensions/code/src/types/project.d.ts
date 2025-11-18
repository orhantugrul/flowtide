export type Project = {
  id: number;
  name: string;
  path: string;
  createdAt?: string;
  updatedAt?: string;
  deletedAt?: string;
};

export type ProjectQueryInput = {
  name?: string;
  path?: string;
};

export type ProjectCreateInput = {
  name: string;
  path: string;
};
