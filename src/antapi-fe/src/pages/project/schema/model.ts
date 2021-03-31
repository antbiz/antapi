import type { Effect, Reducer } from 'umi';
import { getSchemas } from '@/services/schema';

export interface SchemaState {
  data: any;
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
    data: {},
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
        data: action.payload,
      }
    },
    clean(state) {
      return {
        ...(state as SchemaState),
        data: {},
      }
    }
  },
}

export default SchemaModel;
