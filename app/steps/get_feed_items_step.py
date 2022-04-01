import feedparser
from watergrid.context import DataContext, OutputMode
from watergrid.steps import Step


class GetFeedItemsStep(Step):
    def __init__(self):
        super().__init__(self.__class__.__name__, requires=['source'], provides=['item'])

    def run(self, context: DataContext):
        context.set_output_mode(OutputMode.SPLIT)
        feed = feedparser.parse(context.get('source').get_url())
        item_containers = []
        for item in feed.entries:
            item_containers.append(item)
        context.set('item', item_containers)
