import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

const routes: Routes = [
    {
        path: 'notifications',
        loadChildren: () => import('../setting-notifications/setting-notifications.module').then(
            m => m.SettingNotificationsModule
        ),
    },
    {
        path: 'billing',
        loadChildren: () => import('../setting-billing/setting-billing.module').then(
            m => m.SettingBillingModule
        ),
    },
];


@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class SettingShellRoutingModule { }
