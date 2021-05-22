import { createContext, useContext } from 'react';

export type SourceCtxType = {
  currentSchema: API.Schema;
  setCurrentSchema: (schema: API.Schema) => void;
};

export const SourceCtx = createContext<SourceCtxType>({});

export const useSourceCtx = () => useContext<SourceCtxType>(SourceCtx);
