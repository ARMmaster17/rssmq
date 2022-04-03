from watergrid.pipelines import StandalonePipeline
from watergrid.pipelines.pipeline import Pipeline

from app.metrics.console_exporter import ConsoleExporter
from app.steps.get_feed_items_step import GetFeedItemsStep
from app.steps.get_sources_step import GetSourcesStep


def build_pipeline() -> Pipeline:
    pipeline = StandalonePipeline('rssmq_pipeline')
    pipeline.add_metrics_exporter(ConsoleExporter())
    sources = ["test/rss_sample.xml"]
    pipeline.add_step(GetSourcesStep(sources))
    pipeline.add_step(GetFeedItemsStep())
    return pipeline


class App:
    def __init__(self):
        self._pipeline = build_pipeline()

    def run(self):
        self._pipeline.run_loop()
