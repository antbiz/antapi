/**
 * 将fr schema转成 pro-table 的 columns
 */
export const frSchema2ProTableCols = (schema) => {
  const props = schema.properties;
  return Object.keys(props).map(
    field => (
      {
        title: props[field].title,
        dataIndex: field,
      }
    )
  )
};
