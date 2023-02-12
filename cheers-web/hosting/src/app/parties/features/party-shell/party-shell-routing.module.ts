import {NgModule} from '@angular/core';
import {RouterModule, Routes} from '@angular/router';

const routes: Routes = [
    {
        path: 'feed',
        loadChildren: () => import('../party-feed/party-feed.module').then(m => m.PartyFeedModule)
    },
    {
        path: 'hosting',
        loadChildren: () => import('../party-hosting/party-hosting.module').then(m => m.PartyHostingModule)
    },
    {
        path: 'create',
        loadChildren: () => import('../party-create/party-create.module').then(m => m.PartyCreateModule)
    },
    {
        path: ':id/edit',
        loadChildren: () => import('../party-edit/party-edit.module').then(m => m.PartyEditModule)
    },
    {
        path: ':id/guest-list',
        loadChildren: () => import('../party-guest-list/party-guest-list.module').then(m => m.PartyGuestListModule)
    },
];

@NgModule({
    imports: [RouterModule.forChild(routes)],
    exports: [RouterModule]
})
export class PartyShellRoutingModule {
}
