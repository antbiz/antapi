import { Request, Response } from 'express';

export default {
  'GET /api/biz/test_a': {
    data: [
      {
        id: '1',
        username: 'bean',
        phone: '666666',
      },
      {
        id: '2',
        username: 'dou',
        phone: '888888',
      },
    ],
  },
  'GET /api/biz/test_a/1': {
    data: {
      id: '1',
      username: 'bean',
      phone: '666666',
    },
  },
  'GET /api/biz/test_a/2': {
    data: {
      id: '2',
      username: 'dou',
      phone: '888888',
    },
  },
  'GET /api/biz/test_b': {
    data: [
      {
        id: '1',
        title: 'this is a test',
        desc: '这是一个测试',
      },
      {
        id: '2',
        title: 'this is a tttttest',
        desc: '这是一个测测测测试',
      },
    ],
  },
  'GET /api/biz/test_b/1': {
    data: {
      id: '1',
      title: 'this is a test',
      desc: '这是一个测试',
    },
  },
  'GET /api/biz/test_b/2': {
    data: {
      id: '2',
      title: 'this is a tttttest',
      desc: '这是一个测测测测试',
    },
  },
};
