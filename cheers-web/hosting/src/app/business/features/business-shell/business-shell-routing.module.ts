import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

const routes: Routes = [
  {
    path: 'parties',
    loadChildren: () => import('../business-parties/business-parties.module').then(m => m.BusinessPartiesModule)
  },
  {
    path: 'payouts',
    loadChildren: () => import('../business-payouts/business-payouts.module').then(m => m.BusinessPayoutsModule)
  },
  {
    path: 'orders',
    loadChildren: () => import('../business-orders/business-orders.module').then(m => m.BusinessOrdersModule)
  },
  {
    path: 'reports',
    loadChildren: () => import('../business-reports/business-reports.module').then(m => m.BusinessReportsModule)
  },
];
@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class BusinessShellRoutingModule { }
