import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { ChatsRoutingModule } from './chats-routing.module';
import {ChatComponent} from "./feature/chat.component";
import {CoreModule} from "../core/core.module";
import {AppModule} from "../app.module";
import {RoomItemComponent} from "./ui/room-item/room-item.component";
import { ChatContentComponent } from './ui/chat-content/chat-content.component';
import {SharedModule} from "../shared/shared.module";
import {MaterialModule} from "../material/material.module";
import {RouterModule} from "@angular/router";
import {UserChipModule} from "../users/ui/user-chip/user-chip.module";


@NgModule({
  declarations: [
      ChatComponent,
      RoomItemComponent,
      ChatContentComponent,
  ],
  imports: [
      CommonModule,
      ChatsRoutingModule,
      RouterModule,
      MaterialModule,
      UserChipModule,
  ]
})
export class ChatsModule { }
