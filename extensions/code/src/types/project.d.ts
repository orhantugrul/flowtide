export interface Project {
  id?: number;
  name: string;
  path: string;
  createdAt?: Date;
  updatedAt?: Date;
  deletedAt?: Date;
}

export interface ProjectCreateInput {
  name: string;
  path: string;
}
