import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { ChatMessageItemRoutingModule } from './chat-message-item-routing.module';
import { ChatMessageItemComponent } from './chat-message-item.component';
import {ExtendedModule, FlexModule} from "@angular/flex-layout";
import {MatIconModule} from "@angular/material/icon";
import {MatLegacyTooltipModule as MatTooltipModule} from "@angular/material/legacy-tooltip";


@NgModule({
    declarations: [
        ChatMessageItemComponent
    ],
    exports: [
        ChatMessageItemComponent
    ],
    imports: [
        CommonModule,
        ChatMessageItemRoutingModule,
        FlexModule,
        ExtendedModule,
        MatIconModule,
        MatTooltipModule
    ]
})
export class ChatMessageItemModule { }
