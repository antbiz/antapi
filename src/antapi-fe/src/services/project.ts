// @ts-ignore
/* eslint-disable */
import { request } from 'umi';

/** 获取项目列表 GET /api/admin/project */
export async function getProjects(
  params?: {
    // query
    /** 当前的页码 */
    current?: number;
    /** 页面的容量 */
    pageSize?: number;
  },
  options?: { [key: string]: any },
) {
  return request<{ data: API.Project[] }>('/api/admin/project', {
    method: 'GET',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}

/** 获取项目详情  GET /api/project/${projectId} */
export async function getProject(projectId: string, options?: { [key: string]: any }) {
  return request<{ data: API.Project }>(`/api/admin/project/${projectId}`, {
    method: 'GET',
    ...(options || {}),
  });
}

/** 更新项目 PUT /api/admin/project */
export async function updateProject(projectId: string, options?: { [key: string]: any }) {
  return request<{ data: API.Project }>(`/api/admin/project/${projectId}`, {
    method: 'PUT',
    ...(options || {}),
  });
}

/** 新建项目 POST /api/project */
export async function createProject(options?: { [key: string]: any }) {
  return request<{ data: API.Project }>('/api/admin/project', {
    method: 'POST',
    ...(options || {}),
  });
}

/** 删除项目 DELETE /api/project */
export async function deleteProject(projectId: string, options?: { [key: string]: any }) {
  return request(`/api/admin/project/${projectId}`, {
    method: 'DELETE',
    ...(options || {}),
  });
}
