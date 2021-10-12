import { mount } from '@cypress/vue'
import FeedItem from "@/components/FeedItem";

it ('renders feed item', () => {
    mount(FeedItem, {
        propsData: {
            item: {
                ID: 0,
                Url: "example.com/rss.xml",
                LastChecked: "0000 00:00"
            }
        }
    })

    cy.contains('Delete')
})