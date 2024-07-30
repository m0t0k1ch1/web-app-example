import { Injectable, inject } from '@angular/core';

import { NotificationService } from './notification.service';

@Injectable({
  providedIn: 'root',
})
export class ErrorService {
  private notificationService = inject(NotificationService);

  constructor() {}

  public handle(err: any): void {
    this.notificationService.unexpectedError(err);
  }
}
