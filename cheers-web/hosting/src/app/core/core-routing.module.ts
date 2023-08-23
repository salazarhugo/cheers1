import {NgModule} from '@angular/core';
import {RouterModule, Routes} from '@angular/router';
import {ChatComponent} from "../chats/feature/chat.component";
import {HomeComponent} from "../home/features/home/home.component";
import {MapComponent} from "../map/map.component";
import {TicketsComponent} from "../shared/ui/tickets/tickets.component";
import {PartyResolver} from "../manage-parties/data/party.resolver";
import {ProfileResolver} from "../users/data/profile.resolver";

const routes: Routes = [
    {
        path: 'home',
        component: HomeComponent,
    },
    {
        path: 'profile',
        loadChildren: () => import('../users/feature/profile/profile.module').then(m => m.ProfileModule),
    },
    {
        path: 'p/:id',
        resolve: {party: PartyResolver},
        loadChildren: () => import('../parties/features/party-detail/party-detail.module').then(m => m.PartyDetailModule)
    },
    {
        path: 'u/:username',
        resolve: {user: ProfileResolver},
        loadChildren: () => import('../users/feature/user-profile/user-profile.module').then(m => m.UserProfileModule),
    },
    {
        path: 'parties',
        loadChildren: () => import('../parties/features/party-feed/party-feed.module').then(m => m.PartyFeedModule)
    },
    {
        path: 'tickets',
        component: TicketsComponent,
    },
    {
        path: 'friends',
        component: ChatComponent,
    },
    {
        path: 'events',
        redirectTo: 'parties',
    },
    {
        path: 'map',
        component: MapComponent,
    },
    {
        path: 'notifications',
        loadChildren: () => import('../activities/feature/activity-shell/activity-shell.module').then(m => m.ActivityShellModule)
    },
    {
        path: 'chats',
        component: ChatComponent,
        loadChildren: () => import('../chats/chats.module').then(m => m.ChatsModule)
    },
    {
        path: 'parties',
        loadChildren: () => import('../parties/features/party-shell/party-shell.module').then(m => m.PartyShellModule)
    },
    {
        path: 'business',
        loadChildren: () => import('../business/features/business-shell/business-shell.module').then(m => m.BusinessShellModule)
    },
    {
        path: 'settings',
        loadChildren: () => import('../settings/feature/setting-shell/setting-shell.module').then(m => m.SettingShellModule)
    },
    {
        path: 'u',
        loadChildren: () => import('../users/feature/user-shell/user-shell.module').then(m => m.UserShellModule)
    },
    {path: '**', redirectTo: 'home'},
];

@NgModule({
    declarations: [],
    imports: [RouterModule.forChild(routes)],
    exports: [RouterModule]
})
export class CoreRoutingModule {
}
