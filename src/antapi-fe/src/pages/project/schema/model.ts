import type { Effect, Reducer } from 'umi';
import { getSchemas } from '@/services/schema';

export interface SchemaState {
  schemas: Partial<API.Schema[]>;
}

export interface SchemaModelType {
  namespace: string;
  state: SchemaState;
  effects: {
    getSchemas: Effect;

  };
  reducers: {
    loadSchemas: Reducer<SchemaState>;
    clean: Reducer<SchemaState>;
  };
}

const SchemaModel: SchemaModelType = {
  namespace: 'schema',

  state: {
    schemas: [],
  },

  effects: {
    *getSchemas({ payload }, { call, put }) {
      const { projectId } = payload;
      const response = yield call(getSchemas, projectId);
      yield put({
        type: 'loadSchemas',
        payload: response,
      });
    },
  },

  reducers: {
    loadSchemas(state, action) {
      return {
        ...(state as SchemaState),
        // schemas: action.payload?.data?.list || [],
        schemas: action.payload?.data || [],
      }
    },
    clean(state) {
      return {
        ...(state as SchemaState),
        schemas: [],
      }
    }
  },
}

export default SchemaModel;
