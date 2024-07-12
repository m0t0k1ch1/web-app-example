import type { CodegenConfig } from '@graphql-codegen/cli';

const config: CodegenConfig = {
  overwrite: true,
  schema: './schema/gql/*.gql',
  documents: './gql/*.gql',
  generates: {
    './src/gen/graphql-codegen/schema.ts': {
      plugins: [
        'typescript',
        'typescript-operations',
        'typescript-apollo-angular',
      ],
      config: {
        scalars: {
          Int32: 'number',
          Int64: 'string',
          Uint32: 'number',
          Uint64: 'string',
        },
      },
    },
  },
};

export default config;
