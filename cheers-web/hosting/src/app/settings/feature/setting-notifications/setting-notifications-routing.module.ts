import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import {SettingNotificationsComponent} from "./setting-notifications.component";

const routes: Routes = [
    {
        path: '',
        component: SettingNotificationsComponent,
    },
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class SettingNotificationsRoutingModule { }
