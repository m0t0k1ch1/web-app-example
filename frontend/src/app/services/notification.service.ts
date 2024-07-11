import { Injectable } from '@angular/core';

import * as utils from '../utils';

@Injectable({
  providedIn: 'root',
})
export class NotificationService {
  private readonly TOAST_OVERLAY_ELEMENT_ID = 'toast-overlay';

  constructor() {}

  public notifyUnexpectedError(err: any): void {
    const toastOverlay = document.getElementById(this.TOAST_OVERLAY_ELEMENT_ID);
    if (toastOverlay === null) {
      console.error(err);
      return;
    }

    const toast = document.createElement('div');
    {
      toast.className =
        'flex flex-col gap-2 rounded-xl bg-red-500 p-4 text-white';

      const upper = document.createElement('div');
      {
        upper.className = 'flex flex-row justify-between';

        const title = document.createElement('p');
        {
          title.className = 'font-bold';
          title.innerText = 'ERROR';
        }

        const closeButton = document.createElement('button');
        {
          closeButton.className = 'w-6 cursor-pointer';
          closeButton.innerText = '×';
          closeButton.onclick = () => {
            toast.remove();
          };
        }

        upper.appendChild(title);
        upper.appendChild(closeButton);
      }

      const lower = document.createElement('div');
      {
        const description = document.createElement('p');
        {
          description.innerText = utils.stringifyError(err);
        }

        lower.appendChild(description);
      }

      toast.appendChild(upper);
      toast.appendChild(lower);
    }

    toastOverlay.appendChild(toast);
  }
}
