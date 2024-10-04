import { Component, EventEmitter, Output } from '@angular/core';

import { RippleModule } from 'primeng/ripple';

import { AddTaskFormDialogComponent } from '../add-task-form-dialog/add-task-form-dialog.component';

@Component({
  selector: 'page-add-task-button',
  standalone: true,
  imports: [RippleModule, AddTaskFormDialogComponent],
  templateUrl: './add-task-button.component.html',
  styleUrl: './add-task-button.component.css',
})
export class AddTaskButtonComponent {
  @Output() onComplete = new EventEmitter<void>();

  public isHovered = false;
  public isFormDialogVisible = false;
}
