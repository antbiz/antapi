// @ts-ignore
/* eslint-disable */
import { request } from 'umi';

/** 获取内容模型列表 GET `/api/project/${projectId}/schema` */
export async function getSchemas(
  params: {
    // query
    /** 当前的页码 */
    current?: number;
    /** 页面的容量 */
    pageSize?: number;
    projectId: string;
  },
  options?: { [key: string]: any },
) {
  return request<{ data: API.Schema[] }>(`/api/project/${params.projectId}/schema`, {
    method: 'GET',
  });
}

/** 获取内容模型详情  GET `/api/project/${projectId}/schema/${schemaId}` */
export async function getSchema(projectId: string, schemaId: string) {
  return request<{ data: API.Schema }>(`/api/project/${projectId}/schema/${schemaId}`, {
    method: 'GET',
  });
}

/** 新建内容模型 PUT `/api/project/${projectId}/schema/${schemaId}`*/
export async function updateSchema(
  projectId: string,
  schemaId: string,
  options?: { [key: string]: any },
) {
  return request<{ data: API.Schema }>(`/api/project/${projectId}/schema/${schemaId}`, {
    method: 'PUT',
    ...(options || {}),
  });
}

/** 新建内容模型 POST `/api/project/${projectId}/schema/${schemaId}`*/
export async function createSchema(
  projectId: string,
  schemaId: string,
  options?: { [key: string]: any },
) {
  return request<{ data: API.Schema }>(`/api/project/${projectId}/schema/${schemaId}`, {
    method: 'POST',
    ...(options || {}),
  });
}

/** 删除内容模型 DELETE`/api/project/${projectId}/schema/${schemaId}`*/
export async function deleteSchema(
  projectId: string,
  schemaId: string,
  options?: { [key: string]: any },
) {
  return request(`/api/project/${projectId}/schema/${schemaId}`, {
    method: 'DELETE',
    ...(options || {}),
  });
}
