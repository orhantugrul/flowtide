export interface Activity {
  id: number;
  projectId: number;
  editorId: number;
  language: string;
  filePath: string;
  startTime: string;
  endTime: string;
  createdAt?: string;
  updatedAt?: string;
  deletedAt?: string;
}

export interface ActivityCreateInput {
  projectId: number;
  editorId: number;
  language: string;
  filePath: string;
  startTime: string;
  endTime: string;
}
