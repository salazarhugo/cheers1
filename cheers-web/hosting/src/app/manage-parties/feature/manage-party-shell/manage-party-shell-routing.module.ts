import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

const routes: Routes = [
  {
    path: ':id/dashboard',
    loadChildren: () => import('../manage-party-dashboard/manage-party-dashboard.module').then(m => m.ManagePartyDashboardModule)
  },
  {
    path: ':id/guest-list',
    loadChildren: () => import('../manage-party-guest-list/manage-party-guest-list.module').then(m => m.ManagePartyGuestListModule)
  },
  {
    path: ':id/basicinfo',
    loadChildren: () => import('../manage-party-basicinfo/manage-party-basicinfo.module').then(m => m.ManagePartyBasicinfoModule)
  },
  {
    path: ':id/details',
    loadChildren: () => import('../manage-party-details/manage-party-details.module').then(m => m.ManagePartyDetailsModule)
  },
  {
    path: ':id/tickets',
    loadChildren: () => import('../manage-party-tickets/manage-party-tickets.module').then(m => m.ManagePartyTicketsModule)
  },
  {
    path: ':id',
    redirectTo: ':id/dashboard'
  },
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class ManagePartyShellRoutingModule { }
