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
    component: '../layouts/ProjectLayout',
    layout: false,
    routes: [
      {
        path: '/project/:projectId/overview',
        component: './project/overview',
      },
      {
        path: '/project/:projectId/schema',
        component: './project/schema',
      },
      {
        path: '/project/:projectId/content',
        component: './project/content',
      },
      {
        path: '/project/:projectId/setting',
        component: './project/setting',
      },
    ]
  },
  {
    component: './404',
  },
];
