import { Component, EventEmitter, Output } from '@angular/core';
import { CommonModule } from '@angular/common';

import { RippleModule } from 'primeng/ripple';

import { AddTaskFormDialogComponent } from '../add-task-form-dialog/add-task-form-dialog.component';

@Component({
  selector: 'app-add-task-button',
  standalone: true,
  imports: [CommonModule, RippleModule, AddTaskFormDialogComponent],
  templateUrl: './add-task-button.component.html',
  styleUrl: './add-task-button.component.css',
})
export class AddTaskButtonComponent {
  @Output() complete = new EventEmitter<void>();

  public isHovered = false;
  public isFormDialogVisible = false;

  constructor() {}

  public onClick(): void {
    this.isFormDialogVisible = true;
  }
}
