import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { PostDeleteDialogRoutingModule } from './post-delete-dialog-routing.module';
import { PostDeleteDialogComponent } from './post-delete-dialog.component';
import {MatDialogModule} from "@angular/material/dialog";
import {MatButtonModule} from "@angular/material/button";


@NgModule({
  declarations: [
    PostDeleteDialogComponent
  ],
  imports: [
    CommonModule,
    PostDeleteDialogRoutingModule,
    MatDialogModule,
    MatButtonModule
  ],
  exports: [
    PostDeleteDialogComponent
  ],
})
export class PostDeleteDialogModule { }
