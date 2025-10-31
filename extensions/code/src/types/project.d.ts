export interface Project {
  id: number;
  name: string;
  path: string;
  createdAt?: string;
  updatedAt?: string;
  deletedAt?: string;
}

export interface ProjectQueryInput {
  name?: string;
  path?: string;
}

export interface ProjectCreateInput {
  name: string;
  path: string;
}
