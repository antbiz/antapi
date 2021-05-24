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
    name: 'account',
    path: '/account',
    hideInMenu: true,
    menuRender: false,
    routes: [
      {
        name: 'settings',
        path: '/account/settings',
        component: './account/settings',
      }
    ]
  },
  {
    name: 'overview',
    icon: 'eye',
    path: '/project/:projectName/overview',
    component: './project/overview',
  },
  {
    name: 'source',
    icon: 'database',
    path: '/project/:projectName/source',
    component: './project/source',
  },
  {
    name: 'setting',
    icon: 'setting',
    path: '/project/:projectName/setting',
    component: './project/setting',
  },
  {
    component: './404',
  },
];
