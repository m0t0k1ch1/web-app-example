import { ApplicationConfig, inject } from '@angular/core';

import { ApolloClientOptions, InMemoryCache } from '@apollo/client/core';
import { relayStylePagination } from '@apollo/client/utilities';
import { Apollo, APOLLO_OPTIONS } from 'apollo-angular';
import { HttpLink } from 'apollo-angular/http';

export function apolloOptionsFactory(): ApolloClientOptions<any> {
  const httpLink = inject(HttpLink);

  return {
    link: httpLink.create({ uri: 'http://localhost:8080/graphql' }),
    cache: new InMemoryCache({
      resultCaching: true,
      typePolicies: {
        Query: {
          fields: {
            tasks: relayStylePagination(),
          },
        },
      },
    }),
  };
}

export const graphqlProvider: ApplicationConfig['providers'] = [
  Apollo,
  {
    provide: APOLLO_OPTIONS,
    useFactory: apolloOptionsFactory,
  },
];
