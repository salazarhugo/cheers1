import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';

import {PostItemRoutingModule} from './post-item-routing.module';
import {PostItemComponent} from "./post-item.component";
import {MatDividerModule} from "@angular/material/divider";
import {MatIconModule} from "@angular/material/icon";
import {FormsModule} from "@angular/forms";
import {MatCardModule} from "@angular/material/card";
import {FlexModule} from "@angular/flex-layout";
import {MatButtonModule} from "@angular/material/button";
import {MatMenuModule} from "@angular/material/menu";
import {PostDeleteDialogModule} from "../post-delete-dialog/post-delete-dialog.module";


@NgModule({
  declarations: [PostItemComponent],
  imports: [
    CommonModule,
    PostItemRoutingModule,
    MatDividerModule,
    MatIconModule,
    FormsModule,
    MatCardModule,
    FlexModule,
    MatButtonModule,
    MatMenuModule,
      PostDeleteDialogModule,
  ],
  exports: [PostItemComponent],
})
export class PostItemModule { }
