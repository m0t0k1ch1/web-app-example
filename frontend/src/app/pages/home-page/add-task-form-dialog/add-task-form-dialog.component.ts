import { Component, EventEmitter, Input, Output, inject } from '@angular/core';
import { CommonModule } from '@angular/common';
import {
  AbstractControl,
  FormControl,
  FormGroup,
  ReactiveFormsModule,
  Validators,
} from '@angular/forms';
import { firstValueFrom } from 'rxjs';

import { MutationResult } from 'apollo-angular';

import { ButtonModule } from 'primeng/button';
import { DialogModule } from 'primeng/dialog';
import { InputTextModule } from 'primeng/inputtext';

import {
  HomePage_CreateTaskGQL,
  HomePage_CreateTaskMutation,
} from '../../../../gen/graphql-codegen/schema';

import { ErrorService } from '../../../services/error.service';
import { NotificationService } from '../../../services/notification.service';

@Component({
  selector: 'page-add-task-form-dialog',
  standalone: true,
  imports: [
    CommonModule,
    ReactiveFormsModule,
    ButtonModule,
    DialogModule,
    InputTextModule,
  ],
  templateUrl: './add-task-form-dialog.component.html',
  styleUrl: './add-task-form-dialog.component.css',
})
export class AddTaskFormDialogComponent {
  @Input() isVisible!: boolean;
  @Output() isVisibleChange = new EventEmitter<boolean>();
  @Output() onComplete = new EventEmitter<void>();

  private createTaskGQL = inject(HomePage_CreateTaskGQL);

  private errorService = inject(ErrorService);
  private notificationService = inject(NotificationService);

  public form = new FormGroup({
    title: new FormControl('', [Validators.required, Validators.maxLength(32)]),
  });

  public get titleControl(): AbstractControl {
    return this.form.get('title')!;
  }

  public get shouldShowTitleError(): boolean {
    return (
      this.titleControl.invalid &&
      (this.titleControl.dirty || this.titleControl.touched)
    );
  }

  public get titleError(): string | null {
    if (this.titleControl.hasError('required')) {
      return 'required';
    } else if (this.titleControl.hasError('maxlength')) {
      return `must be ${this.titleControl.getError('maxlength').requiredLength} characters or less`;
    }

    return null;
  }

  public async onSubmit(): Promise<void> {
    if (this.form.invalid) {
      this.form.markAllAsTouched();
      return;
    }

    const title = this.titleControl.value;

    {
      let result: MutationResult<HomePage_CreateTaskMutation>;
      {
        try {
          result = await firstValueFrom(
            this.createTaskGQL.mutate({
              input: {
                title: title,
              },
            }),
          );
        } catch (e) {
          this.errorService.handle(e);
          return;
        }
      }

      const err = result.data!.createTask.error;
      if (err !== undefined && err !== null) {
        switch (err.__typename) {
          case 'BadRequestError':
            this.notificationService.badRequest(err.message);
            break;
          default:
            this.errorService.handle(new Error(err.message));
        }
        return;
      }
    }

    this.onComplete.emit();
    this.hide();
  }

  public onCancel(): void {
    this.hide();
  }

  private hide(): void {
    this.form.reset();
    this.isVisible = false;
    this.isVisibleChange.emit(false);
  }
}
