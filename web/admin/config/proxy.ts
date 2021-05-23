/**
 * 在生产环境 代理是无法生效的，所以这里没有生产环境的配置
 * The agent cannot take effect in the production environment
 * so there is no configuration of the production environment
 * For details, please see
 * https://pro.ant.design/docs/deploy
 */
export default {
  dev: {
    'http://localhost:8001/api': {
      target: 'http://localhost:8199',
      changeOrigin: true,
      pathRewrite: { '^http://localhost:8001': '' },
    },
  },
  test: {
    '/api': {
      target: 'http://localhost:8199',
      changeOrigin: true,
      pathRewrite: { '^/api': '' },
    },
  },
  pre: {
    '/api': {
      target: 'your pre url',
      changeOrigin: true,
      pathRewrite: { '^/api': '' },
    },
  },
};
