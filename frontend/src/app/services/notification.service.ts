import { Injectable, inject } from '@angular/core';

import { MessageService } from 'primeng/api';

import * as utils from '../utils';

@Injectable({
  providedIn: 'root',
})
export class NotificationService {
  private messageService = inject(MessageService);

  constructor() {}

  public unexpectedError(err: any): void {
    this.messageService.add({
      severity: 'error',
      summary: 'UNEXPECTED ERROR',
      detail: utils.stringifyError(err),
      life: 5_000,
    });
  }
}
