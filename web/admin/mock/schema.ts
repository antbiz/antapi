import { Request, Response } from 'express';

const fakeSchemas = [
  {
    type: 'object',
    properties: {
      username: {
        title: '姓名',
        type: 'string',
      },
      phone: {
        title: '电话',
        type: 'string',
      },
    },
    column: 2,
    displayType: 'column',
    showDescIcon: true,
    id: '1',
    createdAt: '2021-03-28',
    updatedAt: '2021-03-28',
    projectName: 'example1',
    collectionName: 'test_a',
    title: '测试A',
    description: '这是一个测试'
  },
  {
    type: 'object',
    properties: {
      title: {
        title: '标题',
        type: 'string',
        'options': {},
      },
      desc: {
        title: '描述',
        type: 'string',
        'options': {},
      },
    },
    column: 2,
    displayType: 'column',
    showDescIcon: true,
    id: '2',
    createdAt: '2021-03-28',
    updatedAt: '2021-03-28',
    projectName: 'example1',
    collectionName: 'test_b',
    title: '测试B',
    description: '这是一个测试'
  },
]

export default {
  'GET /api/biz/schema': {
    data: fakeSchemas,
  },
  'GET /api/biz/schema/1': {
    data: fakeSchemas[0],
  },
  'GET /api/biz/schema/2': {
    data: fakeSchemas[1],
  },
};
