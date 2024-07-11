import { Component, EventEmitter, Input, Output } from '@angular/core';
import { CommonModule } from '@angular/common';
import {
  AbstractControl,
  FormControl,
  FormGroup,
  ReactiveFormsModule,
  Validators,
} from '@angular/forms';

import { ButtonModule } from 'primeng/button';
import { DialogModule } from 'primeng/dialog';
import { InputTextModule } from 'primeng/inputtext';

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

  public onSubmit(): void {
    if (this.form.invalid) {
      return;
    }

    // TODO: add task

    this.hide();
  }

  public onCancel(): void {
    this.hide();
  }

  public hide(): void {
    this.form.reset();
    this.isVisible = false;
    this.isVisibleChange.emit(false);
  }
}
