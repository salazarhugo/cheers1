import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';

import {PostItemRoutingModule} from './post-item-routing.module';
import {PostItemComponent} from "./post-item.component";
import {MatDividerModule} from "@angular/material/divider";
import {MatIconModule} from "@angular/material/icon";
import {FormsModule} from "@angular/forms";
import {MatLegacyCardModule as MatCardModule} from "@angular/material/legacy-card";
import {FlexModule} from "@angular/flex-layout";
import {MatLegacyButtonModule as MatButtonModule} from "@angular/material/legacy-button";
import {MatLegacyMenuModule as MatMenuModule} from "@angular/material/legacy-menu";
import {PostDeleteDialogModule} from "../post-delete-dialog/post-delete-dialog.module";
import {RelativeTimeModule} from "../../../shared/data/pipes/relative-time/relative-time.module";
import {MatLegacyTooltipModule as MatTooltipModule} from "@angular/material/legacy-tooltip";


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
        RelativeTimeModule,
        MatTooltipModule,
    ],
  exports: [PostItemComponent],
})
export class PostItemModule { }
