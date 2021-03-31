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
    name: 'overview',
    path: '/project/:projectId/overview',
    icon: 'eye',
    component: './project/overview',
  },
  {
    name: 'schema',
    path: '/project/:projectId/schema',
    icon: 'gold',
    component: './project/schema',
  },
  {
    name: 'content',
    path: '/project/:projectId/content',
    icon: 'database',
    component: './project/content',
  },
  {
    name: 'setting',
    path: '/project/:projectId/setting',
    icon: 'setting',
    component: './project/setting',
  },
  {
    component: './404',
  },
];
