import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { ChatNewRoutingModule } from './chat-new-routing.module';
import { ChatNewComponent } from './chat-new.component';
import {MatButtonModule} from "@angular/material/button";
import {MatIconModule} from "@angular/material/icon";
import {FlexModule} from "@angular/flex-layout";


@NgModule({
  declarations: [
    ChatNewComponent
  ],
    imports: [
        CommonModule,
        ChatNewRoutingModule,
        MatButtonModule,
        MatIconModule,
        FlexModule
    ],
  exports: [
    ChatNewComponent
  ],
})
export class ChatNewModule { }
