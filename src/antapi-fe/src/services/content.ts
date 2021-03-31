// @ts-ignore
/* eslint-disable */
import { request } from 'umi';

export interface Options {
  page?: number;
  pageSize?: number;

  filter?: {
    id?: string;
    ids?: string[];
    [key: string]: any;
  };

  fuzzyFilter?: {
    [key: string]: any;
  };

  sort?: {
    [key: string]: 'ascend' | 'descend' | null;
  };

  payload?: Record<string, any>;
}

export async function getContents(projectId: string, resource: string, options?: Options) {
  return request(`/api/project/${projectId}/content`, {
    method: 'GET',
    data: {
      options,
      resource,
      action: 'getMany',
    },
  });
}

export async function getContent(
  projectId: string,
  contentId: string,
  resource: string,
  options?: Options,
) {
  return request(`/api/project/${projectId}/content/${contentId}`, {
    method: 'GET',
    data: {
      options,
      resource,
      action: 'getMany',
    },
  });
}

export async function createContent(
  projectId: string,
  resource: string,
  payload: Record<string, any>,
) {
  return request(`/api/project/${projectId}/content`, {
    method: 'POST',
    data: {
      resource,
      action: 'createOne',
      options: {
        payload,
      },
    },
  });
}

export async function deleteContent(projectId: string, resource: string, id: string) {
  return request(`/api/project/${projectId}/content`, {
    method: 'POST',
    data: {
      resource,
      options: {
        filter: {
          id: id,
        },
      },
      action: 'deleteOne',
    },
  });
}

export async function batchDeleteContent(projectId: string, resource: string, ids: string[]) {
  return request(`/api/project/${projectId}/content`, {
    method: 'POST',
    data: {
      resource,
      options: {
        filter: {
          ids,
        },
      },
      action: 'deleteMany',
    },
  });
}

/**
 * 更新内容
 */
export async function updateContent(
  projectId: string,
  resource: string,
  id: string,
  payload: Record<string, any>,
) {
  return request(`/api/project/${projectId}/content`, {
    method: 'POST',
    data: {
      resource,
      options: {
        payload,
        filter: {
          id: id,
        },
      },
      action: 'updateOne',
    },
  });
}

export async function getMigrateJobs(projectId: string, page = 1, pageSize = 10) {
  return request(`/api/project/${projectId}/migrate`, {
    method: 'GET',
    params: {
      page,
      pageSize,
    },
  });
}

export async function createMigrateJobs(
  projectId: string,
  collectionName: string,
  filePath: string,
  conflictMode: string,
) {
  return request(`/api/project/${projectId}/migrate`, {
    method: 'POST',
    data: {
      filePath,
      projectId,
      conflictMode,
      collectionName,
    },
  });
}
