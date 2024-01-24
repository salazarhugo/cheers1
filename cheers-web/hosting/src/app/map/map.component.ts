import {Component, OnInit} from '@angular/core';
import * as mapboxgl from 'mapbox-gl';
import {PostService} from "../posts/data/post.service";
import {PostResponse} from "../../gen/ts/cheers/post/v1/post_service";

@Component({
    selector: 'app-map',
    templateUrl: './map.component.html',
    styleUrls: ['./map.component.sass']
})
export class MapComponent implements OnInit {
    map!: mapboxgl.Map;
    style = 'mapbox://styles/mapbox/streets-v11';
    lat = 37.75;
    lng = -122.41;
    posts: PostResponse[]

    constructor(
        private postService: PostService,
    ) {
    }

    ngOnInit(): void {
    }

    onMapLoad(map: mapboxgl.Map) {
        this.postService.listMapPost().subscribe(posts => {
            this.posts = posts
            posts.forEach(post => {
                if (post.post == undefined)
                    return

                const el = document.createElement('div');
                const width = 50;
                const height = 50;
                el.className = 'marker';
                el.style.backgroundImage = `url(${post.post.photos[0]})`;
                el.style.width = `${width}px`;
                el.style.height = `${height}px`;
                el.style.backgroundSize = '100%';

                const marker = new mapboxgl.Marker(el)
                    .setLngLat([post.post.longitude, post.post.latitude])
                    .addTo(map);
            })
        })
    }
}
