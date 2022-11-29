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
import {GrpcCoreModule} from "@ngx-grpc/core";
import {GrpcWebClientModule} from "@ngx-grpc/grpc-web-client";


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
      GrpcCoreModule.forRoot(),
      GrpcWebClientModule.forRoot({
          settings: {host: 'https://chat-r3a2dr4u4a-nw.a.run.app:443'},
      }),
  ]
})
export class ChatsModule { }
