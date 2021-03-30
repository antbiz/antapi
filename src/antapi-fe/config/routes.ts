export default [
  {
    path: '/',
    redirect: '/home',
  },
  {
    name: 'login',
    layout: false,
    path: '/login',
    component: './login',
    hideInMenu: true,
  },
  {
    name: 'home',
    path: '/home',
    component: './home',
    hideInMenu: true,
    menuRender: false,
  },
  {
    component: '../layout',
    layout: false,
    routes: [
      {
        name: 'overview',
        path: '/:projectId/home',
        icon: 'eye',
        component: './project/overview',
      },
      {
        name: 'schema',
        path: '/:projectId/schema',
        icon: 'gold',
        component: './project/schema',
      },
      {
        name: 'content',
        path: '/:projectId/content',
        icon: 'database',
        routes: [
          {
            exact: true,
            path: '/:projectId/content/:schemaId',
            icon: 'database',
            component: './project/content',
          },
          {
            exact: true,
            path: '/:projectId/content/:schemaId/edit',
            component: './project/content/ContentEditor',
          },
        ],
      },
      {
        name: 'setting',
        path: '/:projectId/setting',
        icon: 'setting',
        component: './project/setting',
      },
    ],
  },
  {
    component: './404',
  },
];
