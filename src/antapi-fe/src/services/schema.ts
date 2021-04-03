// @ts-ignore
/* eslint-disable */
import { request } from 'umi';

/** 获取内容模型列表 GET `/api/admin/schema?project_name=${projectName}` */
export async function getSchemas(
  params: {
    projectName: string
  }
) {
  return request<{ data: API.Schema[] }>(`/api/admin/schema?projectName=${params.projectName}`, {
    method: 'GET',
  });
}

/** 获取内容模型详情  GET `/api/admin/schema/${schemaId}` */
export async function getSchema(schemaId: string) {
  return request<{ data: API.Schema }>(`/api/admin/schema/${schemaId}`, {
    method: 'GET',
  });
}

/** 新建内容模型 PUT `/api/admin/schema/${schemaId}`*/
export async function updateSchema(
  schemaId: string,
  options?: { [key: string]: any },
) {
  return request<{ data: API.Schema }>(`/api/admin/schema/${schemaId}`, {
    method: 'PUT',
    ...(options || {}),
  });
}

/** 新建内容模型 POST `/api/admin/schema`*/
export async function createSchema(
  options?: { [key: string]: any },
) {
  return request<{ data: API.Schema }>(`/api/admin/schema`, {
    method: 'POST',
    ...(options || {}),
  });
}

/** 删除内容模型 DELETE `/api/admin/schema/${schemaId}`*/
export async function deleteSchema(
  schemaId: string,
  options?: { [key: string]: any },
) {
  return request(`/api/admin/schema/${schemaId}`, {
    method: 'DELETE',
    ...(options || {}),
  });
}
