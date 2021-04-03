import { Request, Response } from 'express';

const fakeSchemas = [
  {
    schema: {
      type: 'object',
      properties: {
        username: {
          title: '姓名',
          type: 'string',
          'ui:options': {},
        },
        phone: {
          title: '电话',
          type: 'string',
          'ui:options': {},
        },
      },
      'ui:displayType': 'row',
      'ui:showDescIcon': true,
    },
    displayType: 'row',
    showDescIcon: true,
    id: '1',
    createdAt: '2021-03-28',
    updatedAt: '2021-03-28',
    projectName: 'example1',
    name: 'test_a',
    title: '测试A',
    description: '这是一个测试'
  },
  {
    schema: {
      type: 'object',
      properties: {
        title: {
          title: '标题',
          type: 'string',
          'ui:options': {},
        },
        desc: {
          title: '描述',
          type: 'string',
          'ui:options': {},
        },
      },
      'ui:displayType': 'row',
      'ui:showDescIcon': true,
    },
    displayType: 'row',
    showDescIcon: true,
    id: '2',
    createdAt: '2021-03-28',
    updatedAt: '2021-03-28',
    projectName: 'example1',
    name: 'test_b',
    title: '测试B',
    description: '这是一个测试'
  },
]

export default {
  'GET /api/admin/schema': {
    data: fakeSchemas,
  },
  'GET /api/admin/schema/1': {
    data: fakeSchemas[0],
  },
  'GET /api/admin/schema/2': {
    data: fakeSchemas[1],
  },
};
