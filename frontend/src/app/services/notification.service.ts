import { Injectable, inject } from '@angular/core';
import { MatSnackBar } from '@angular/material/snack-bar';

import * as utils from '../utils';

@Injectable({
  providedIn: 'root',
})
export class NotificationService {
  private snackBar = inject(MatSnackBar);

  constructor() {}

  public notifyUnexpectedError(err: any): void {
    this.snackBar.open(`ERROR: ${utils.stringifyError(err)}`, 'Close');
  }
}
