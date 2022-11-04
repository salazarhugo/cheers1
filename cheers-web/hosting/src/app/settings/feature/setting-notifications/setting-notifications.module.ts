import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';

import {SettingNotificationsRoutingModule} from './setting-notifications-routing.module';
import {SettingNotificationsComponent} from "./setting-notifications.component";


@NgModule({
    declarations: [SettingNotificationsComponent],
    imports: [
        CommonModule,
        SettingNotificationsRoutingModule
    ],
    exports: [SettingNotificationsComponent]
})
export class SettingNotificationsModule {
}
