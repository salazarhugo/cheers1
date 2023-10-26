import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { PostDeleteDialogRoutingModule } from './post-delete-dialog-routing.module';
import { PostDeleteDialogComponent } from './post-delete-dialog.component';
import {MatLegacyDialogModule as MatDialogModule} from "@angular/material/legacy-dialog";
import {MatLegacyButtonModule as MatButtonModule} from "@angular/material/legacy-button";


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
