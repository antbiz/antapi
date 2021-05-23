// @ts-ignore
/* eslint-disable */
import { request } from 'umi';

/** 获取项目列表 GET /api/biz/project */
export async function getProjects() {
  return request<{ data: API.Project[] }>('/api/biz/project?filter={"name":{"$ne":"system"}}', {
    method: 'GET',
  });
}

/** 获取项目详情  GET /api/project/${projectId} */
export async function getProject(projectId: string) {
  return request<{ data: API.Project }>(`/api/biz/project/${projectId}`, {
    method: 'GET',
  });
}

/** 获取项目详情  GET /api/project/find */
export async function findProject(params?: { [key: string]: any }) {
  return request<{ data: API.Project }>('/api/biz/project/find', {
    method: 'GET',
    params: params,
  });
}

/** 更新项目 PUT /api/biz/project */
export async function updateProject(projectId: string, data?: { [key: string]: any }) {
  return request<{ data: API.Project }>(`/api/biz/project/${projectId}`, {
    method: 'PUT',
    data: data,
  });
}

/** 新建项目 POST /api/project */
export async function createProject(data?: { [key: string]: any }) {
  return request<{ data: API.Project }>('/api/biz/project', {
    method: 'POST',
    data: data,
  });
}

/** 删除项目 DELETE /api/project */
export async function deleteProject(projectId: string) {
  return request(`/api/biz/project/${projectId}`, {
    method: 'DELETE',
  });
}

/** 获取项目权限 GET `/api/biz/permission?projectName=${projectName}` */
export async function getPermissions(params: { projectName: string }) {
  return request<{ data: API.Permission[] }>(`/api/biz/permission?filter={"projectName":"${params.projectName}"}`, {
    method: 'GET',
  });
}

/** 更新项目权限 PUT `/api/biz/permission/${id}`*/
export async function updatePermission(id: string, data?: { [key: string]: any }) {
  return request<{ data: API.Permission }>(`/api/biz/permission/${id}`, {
    method: 'PUT',
    data: data,
  });
}
