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

import { ButtonModule } from 'primeng/button';
import { DialogModule } from 'primeng/dialog';
import { InputTextModule } from 'primeng/inputtext';

import { CreateTaskGQL } from '../../../gen/graphql-codegen/schema';

import { NotificationService } from '../../services/notification.service';

@Component({
  selector: 'app-add-task-form-dialog',
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
  @Output() complete = new EventEmitter<void>();

  private createTaskGQL = inject(CreateTaskGQL);

  private notificationService = inject(NotificationService);

  public form = new FormGroup({
    title: new FormControl('', [Validators.required, Validators.maxLength(32)]),
  });

  constructor() {}

  public get titleControl(): AbstractControl {
    return this.form.get('title')!;
  }

  public get shouldShowTitleError(): boolean {
    return (
      this.titleControl.invalid &&
      (this.titleControl.dirty || this.titleControl.touched)
    );
  }

  public async onSubmit(): Promise<void> {
    if (this.form.invalid) {
      return;
    }

    const title = this.titleControl.value;

    try {
      await firstValueFrom(
        this.createTaskGQL.mutate({
          title: title,
        }),
      );
    } catch (e) {
      this.notificationService.unexpectedError(e);
      return;
    }

    this.complete.emit();
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
