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
    icon: 'eye',
    path: '/project/:projectId/overview',
    component: './project/overview',
  },
  {
    name: 'source',
    icon: 'database',
    path: '/project/:projectId/source',
    component: './project/source',
  },
  {
    name: 'setting',
    icon: 'setting',
    path: '/project/:projectId/setting',
    component: './project/setting',
  },
  {
    component: './404',
  },
];
