export interface Activity {
  id?: number;
  projectId: number;
  editorId: number;
  language: string;
  filePath: string;
  startTime: Date;
  endTime: Date;
  createdAt?: Date;
  updatedAt?: Date;
  deletedAt?: Date;
}

export interface ActivityCreateInput {
  projectId: number;
  editorId: number;
  language: string;
  filePath: string;
  startTime: Date;
  endTime: Date;
}
