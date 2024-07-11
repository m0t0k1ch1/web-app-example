import { Injectable, inject } from '@angular/core';

import * as utils from '../utils';

@Injectable({
  providedIn: 'root',
})
export class NotificationService {
  constructor() {}

  public notifyUnexpectedError(err: any): void {
    // TODO: improve
    console.log(`ERROR: ${utils.stringifyError(err)}`);
  }
}
