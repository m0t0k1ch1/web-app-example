import { Injectable, inject } from '@angular/core';

import { MessageService } from 'primeng/api';

import * as utils from '../utils';

@Injectable({
  providedIn: 'root',
})
export class NotificationService {
  private primengMessageService = inject(MessageService);

  constructor() {}

  public badRequest(err: any): void {
    this.primengMessageService.add({
      severity: 'error',
      summary: 'BAD REQUEST',
      detail: utils.stringifyError(err),
      life: 5_000,
    });
  }

  public unexpectedError(err: any): void {
    this.primengMessageService.add({
      severity: 'error',
      summary: 'UNEXPECTED ERROR',
      detail: utils.stringifyError(err),
      life: 5_000,
    });
  }
}
